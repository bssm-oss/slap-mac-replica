class SlapMacReplica < Formula
  desc "Play a sound when you physically slap an Apple Silicon MacBook"
  homepage "https://github.com/bssm-oss/slap-mac-replica"
  url "https://github.com/bssm-oss/slap-mac-replica/releases/download/v0.1.4/slap-mac-replica_0.1.4_darwin_arm64.tar.gz"
  sha256 "68865d067a296923dabce873bb2db92eb4f5363d43300b4dc895c13582604f03"
  license "MIT"

  def install
    odie "slap-mac-replica requires an Apple Silicon Mac." if Hardware::CPU.intel?
    bin.install "slap-mac-replica"
    pkgshare.install "presets"
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
