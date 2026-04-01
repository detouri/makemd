class makemd < Formula
  desc "Generate and maintain purposeful README.md files"
  homepage "https://github.com/your-org/makemd"
  version "makemd"

  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/your-org/makemd/releases/download/v#{version}/makemd_darwin_arm64.tar.gz"
      sha256 "REPLACE_ME"
    else
      url "https://github.com/your-org/makemd/releases/download/v#{version}/makemd_darwin_amd64.tar.gz"
      sha256 "REPLACE_ME"
    end
  elsif OS.linux?
    if Hardware::CPU.arm?
      url "https://github.com/your-org/makemd/releases/download/v#{version}/makemd_linux_arm64.tar.gz"
      sha256 "REPLACE_ME"
    else
      url "https://github.com/your-org/makemd/releases/download/v#{version}/makemd_linux_amd64.tar.gz"
      sha256 "REPLACE_ME"
    end
  end

  def install
    bin.install "makemd"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/makemd version")
  end
end