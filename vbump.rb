class Vbump < Formula
  desc "Bump your version and push to git in a safe and predictable manner"
  homepage "https://github.com/calm/homebrew-vbump"
  url "https://github.com/calm/homebrew-vbump/archive/v0.0.8.tar.gz"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/calm/homebrew-vbump"
    system "go", "get", "-u", "github.com/calm/homebrew-vbump"
    cd path do
      system "go", "build", "-o", "#{bin}/vbump"
    end
  end

  test do
    system "#{bin}/vbump", "--help"
  end
end
