package builder

import (
	"errors"
	"fmt"
	"path/filepath"
)

type MetadataBuilder struct {
	formDirectoryReader           metadataPartsDirectoryReader
	instanceGroupDirectoryReader  metadataPartsDirectoryReader
	jobsDirectoryReader           metadataPartsDirectoryReader
	iconEncoder                   iconEncoder
	logger                        logger
	metadataReader                metadataReader
	propertiesDirectoryReader     metadataPartsDirectoryReader
	releaseManifestReader         releaseManifestReader
	runtimeConfigsDirectoryReader metadataPartsDirectoryReader
	stemcellManifestReader        stemcellManifestReader
	variablesDirectoryReader      metadataPartsDirectoryReader
}

type BuildInput struct {
	FormDirectories          []string
	IconPath                 string
	InstanceGroupDirectories []string
	JobDirectories           []string
	MetadataPath             string
	PropertyDirectories      []string
	ReleaseTarballs          []string
	RuntimeConfigDirectories []string
	StemcellTarball          string
	Variables                map[string]string
	BOSHVariableDirectories  []string
	Version                  string
}

//go:generate counterfeiter -o ./fakes/release_manifest_reader.go --fake-name ReleaseManifestReader . releaseManifestReader

type releaseManifestReader interface {
	Read(path string) (ReleaseManifest, error)
}

//go:generate counterfeiter -o ./fakes/stemcell_manifest_reader.go --fake-name StemcellManifestReader . stemcellManifestReader

type stemcellManifestReader interface {
	Read(path string) (StemcellManifest, error)
}

//go:generate counterfeiter -o ./fakes/metadata_parts_directory_reader.go --fake-name MetadataPartsDirectoryReader . metadataPartsDirectoryReader

type metadataPartsDirectoryReader interface {
	Read(path string) ([]Part, error)
}

//go:generate counterfeiter -o ./fakes/metadata_reader.go --fake-name MetadataReader . metadataReader

type metadataReader interface {
	Read(path, version string) (Metadata, error)
}

//go:generate counterfeiter -o ./fakes/icon_encoder.go --fake-name IconEncoder . iconEncoder

type iconEncoder interface {
	Encode(path string) (string, error)
}

type logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

func NewMetadataBuilder(
	formDirectoryReader metadataPartsDirectoryReader,
	instanceGroupDirectoryReader metadataPartsDirectoryReader,
	jobsDirectoryReader metadataPartsDirectoryReader,
	propertiesDirectoryReader metadataPartsDirectoryReader,
	releaseManifestReader releaseManifestReader,
	runtimeConfigsDirectoryReader,
	variablesDirectoryReader metadataPartsDirectoryReader,
	stemcellManifestReader stemcellManifestReader,
	metadataReader metadataReader,
	logger logger,
	iconEncoder iconEncoder,
) MetadataBuilder {
	return MetadataBuilder{
		formDirectoryReader:           formDirectoryReader,
		instanceGroupDirectoryReader:  instanceGroupDirectoryReader,
		jobsDirectoryReader:           jobsDirectoryReader,
		iconEncoder:                   iconEncoder,
		logger:                        logger,
		metadataReader:                metadataReader,
		propertiesDirectoryReader:     propertiesDirectoryReader,
		releaseManifestReader:         releaseManifestReader,
		runtimeConfigsDirectoryReader: runtimeConfigsDirectoryReader,
		stemcellManifestReader:        stemcellManifestReader,
		variablesDirectoryReader:      variablesDirectoryReader,
	}
}

func (m MetadataBuilder) Build(input BuildInput) (GeneratedMetadata, error) {
	metadata, err := m.metadataReader.Read(input.MetadataPath, input.Version)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	productName, ok := metadata["name"].(string)
	if !ok {
		return GeneratedMetadata{}, errors.New("missing \"name\" in tile metadata file")
	}

	releases, err := m.buildReleaseMetadata(input.ReleaseTarballs)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	runtimeConfigs, err := m.buildRuntimeConfigMetadata(input.RuntimeConfigDirectories, metadata)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	variables, err := m.buildVariables(input.BOSHVariableDirectories, metadata)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	m.logger.Printf("Reading stemcell manifest")
	stemcellManifest, err := m.stemcellManifestReader.Read(input.StemcellTarball)
	if err != nil {
		return GeneratedMetadata{}, err
	}
	m.logger.Printf("- version %s", stemcellManifest.Version)

	encodedIcon, err := m.iconEncoder.Encode(input.IconPath)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	formTypes, err := m.buildFormTypes(input.FormDirectories, metadata)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	jobTypes, err := m.buildJobTypes(input.InstanceGroupDirectories, input.JobDirectories, metadata)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	propertyBlueprints, err := m.buildPropertyBlueprints(input.PropertyDirectories, metadata)
	if err != nil {
		return GeneratedMetadata{}, err
	}

	delete(metadata, "name")
	delete(metadata, "icon_image")
	delete(metadata, "form_types")
	delete(metadata, "job_types")
	delete(metadata, "property_blueprints")

	return GeneratedMetadata{
		FormTypes:          formTypes,
		JobTypes:           jobTypes,
		IconImage:          encodedIcon,
		Name:               productName,
		PropertyBlueprints: propertyBlueprints,
		Releases:           releases,
		RuntimeConfigs:     runtimeConfigs,
		Variables:          variables,
		StemcellCriteria: StemcellCriteria{
			OS:          stemcellManifest.OperatingSystem,
			Version:     stemcellManifest.Version,
			RequiresCPI: false,
		},
		Metadata: metadata,
	}, nil
}

func (m MetadataBuilder) buildPropertyBlueprints(dirs []string, metadata Metadata) ([]Part, error) {
	var propertyBlueprints []Part

	if len(dirs) > 0 {
		for _, propertiesDirectory := range dirs {
			m.logger.Printf("Reading property blueprints from %s", propertiesDirectory)

			p, err := m.propertiesDirectoryReader.Read(propertiesDirectory)
			if err != nil {
				return nil,
					fmt.Errorf("error reading from properties directory %q: %s", propertiesDirectory, err)
			}

			propertyBlueprints = append(propertyBlueprints, p...)
		}
	} else {
		if pb, ok := metadata["property_blueprints"].([]interface{}); ok {
			for _, p := range pb {
				propertyBlueprints = append(propertyBlueprints, Part{Metadata: p})
			}
		}
	}

	return propertyBlueprints, nil
}

func (m MetadataBuilder) buildReleaseMetadata(releaseTarballs []string) ([]Release, error) {
	var releases []Release

	m.logger.Printf("Reading release manifests")

	for _, releaseTarball := range releaseTarballs {

		releaseManifest, err := m.releaseManifestReader.Read(releaseTarball)
		if err != nil {
			return nil, err
		}

		m.logger.Printf("- %s", releaseManifest.Name)

		releases = append(releases, Release{
			Name:    releaseManifest.Name,
			Version: releaseManifest.Version,
			File:    filepath.Base(releaseTarball),
		})
	}

	return releases, nil
}

func (m MetadataBuilder) buildRuntimeConfigMetadata(dirs []string, metadata Metadata) ([]Part, error) {
	if _, ok := metadata["runtime_configs"]; ok {
		return nil, errors.New("runtime_config section must be defined using --runtime-configs-directory flag")
	}

	var runtimeConfigs []Part

	for _, runtimeConfigsDirectory := range dirs {
		m.logger.Printf("Reading runtime configs from %s", runtimeConfigsDirectory)

		r, err := m.runtimeConfigsDirectoryReader.Read(runtimeConfigsDirectory)
		if err != nil {
			return nil,
				fmt.Errorf("error reading from runtime configs directory %q: %s", runtimeConfigsDirectory, err)
		}

		runtimeConfigs = append(runtimeConfigs, r...)
	}

	return runtimeConfigs, nil
}

func (m MetadataBuilder) buildVariables(vars []string, metadata Metadata) ([]Part, error) {
	if _, ok := metadata["variables"]; ok {
		return nil, errors.New("variables section must be defined using --variables-directory flag")
	}

	var variables []Part

	for _, variablesDirectory := range vars {
		m.logger.Printf("Reading variables from %s", variablesDirectory)

		v, err := m.variablesDirectoryReader.Read(variablesDirectory)
		if err != nil {
			return nil,
				fmt.Errorf("error reading from variables directory %q: %s", variablesDirectory, err)
		}

		variables = append(variables, v...)
	}

	return variables, nil
}

func (m MetadataBuilder) buildFormTypes(dirs []string, metadata Metadata) ([]Part, error) {
	var formTypes []Part

	if len(dirs) > 0 {
		for _, fd := range dirs {
			m.logger.Printf("Reading forms from %s", fd)

			formTypesInDir, err := m.formDirectoryReader.Read(fd)
			if err != nil {
				return nil, fmt.Errorf("error reading from form directory %q: %s", fd, err)
			}
			formTypes = append(formTypes, formTypesInDir...)
		}
	} else {
		if ft, ok := metadata["form_types"].([]interface{}); ok {
			for _, f := range ft {
				formTypes = append(formTypes, Part{Metadata: f})
			}
		}
	}

	return formTypes, nil
}

func (m MetadataBuilder) buildJobTypes(typeDirs, jobDirs []string, metadata Metadata) ([]Part, error) {
	var jobTypes []Part

	if len(typeDirs) > 0 {
		for _, typeDir := range typeDirs {
			m.logger.Printf("Reading instance groups from %s", typeDir)

			jobTypesInDir, err := m.instanceGroupDirectoryReader.Read(typeDir)
			if err != nil {
				return nil, fmt.Errorf("error reading from instance group directory %q: %s", typeDir, err)
			}
			jobTypes = append(jobTypes, jobTypesInDir...)
		}
	} else {
		if jt, ok := metadata["job_types"].([]interface{}); ok {
			for _, j := range jt {
				jobTypes = append(jobTypes, Part{Metadata: j})
			}
		}
	}

	jobs := map[interface{}]Part{}

	for _, jobDir := range jobDirs {
		m.logger.Printf("Reading jobs from %s", jobDir)

		jobsInDir, err := m.jobsDirectoryReader.Read(jobDir)
		if err != nil {
			return nil, fmt.Errorf("error reading from job directory %q: %s", jobDir, err)
		}
		for _, j := range jobsInDir {
			jobs[j.File] = j
		}
	}

	for _, jt := range jobTypes {
		if templates, ok := jt.Metadata.(map[interface{}]interface{})["templates"]; ok {
			for i, template := range templates.([]interface{}) {
				switch template.(type) {
				case string:
					if job, ok := jobs[template]; ok {
						templates.([]interface{})[i] = job.Metadata
					} else {
						return nil, fmt.Errorf("instance group %q references non-existent job %q",
							jt.Name,
							template,
						)
					}
				}
			}
		}
	}

	return jobTypes, nil
}
