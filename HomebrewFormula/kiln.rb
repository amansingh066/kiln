# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.0.0-dev-pr265.5"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.0.0-dev-pr265.5/kiln-darwin-0.0.0-dev-pr265.5.tar.gz"
      sha256 "72bb99897c2593fb062c0749133a0369c1000b64b12ea835b61044490b120abc"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.0.0-dev-pr265.5/kiln-linux-0.0.0-dev-pr265.5.tar.gz"
      sha256 "f29dbae60841fa46e2b062ae41b20261d83b811df24e9ca6d2bfb5476ae530b2"
    end
  end

  def install
    bin.install "kiln"
  end

  test do
    system "#{bin}/kiln --version"
  end
end
