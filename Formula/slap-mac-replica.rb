class SlapMacReplica < Formula
  desc "Play a sound when you physically slap an Apple Silicon MacBook"
  homepage "https://github.com/bssm-oss/slap-mac-replica"
  url "https://github.com/bssm-oss/slap-mac-replica/releases/download/v0.1.1/slap-mac-replica_0.1.1_darwin_arm64.tar.gz"
  sha256 "9208b8aa42b3541a976998ca002b56b04ca67124d2fc9e9c53cc508f238ed413"
  license "MIT"

  def install
    odie "slap-mac-replica requires an Apple Silicon Mac." if Hardware::CPU.intel?
    bin.install "slap-mac-replica"
  end

  service do
    run [opt_bin/"slap-mac-replica", "run"]
    require_root true
    keep_alive true
    log_path var/"log/slap-mac-replica.log"
    error_log_path var/"log/slap-mac-replica.err.log"
  end

  test do
    output = shell_output("#{bin}/slap-mac-replica doctor")
    assert_match "platform:", output
    assert_match "sensor:", output
  end
end
