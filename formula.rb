class Vbump < Formula
  desc "Bump your version and push to git in a safe and predictable manner"
  homepage "https://github.com/calm/vbump"
  url "https://github.com/calm/vbump/archive/v0.0.4.tar.gz"

  depends_on "go" => :build

  def install
    system "./gobuild.sh"
    bin.install ".gobuild/bin/vbump" => "vbump"
  end

  test do
    system "#{bin}/vbump", "--help"
  end
end
