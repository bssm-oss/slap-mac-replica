class SlapMacReplica < Formula
  desc "Play a sound when you physically slap an Apple Silicon MacBook"
  homepage "https://github.com/bssm-oss/slap-mac-replica"
  url "https://github.com/bssm-oss/slap-mac-replica/releases/download/v0.1.3/slap-mac-replica_0.1.3_darwin_arm64.tar.gz"
  sha256 "790cdf4c757bbc037ce039aa3e463c57e819f66123be2394e3e1810a52242c77"
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
