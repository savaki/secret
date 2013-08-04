#   Copyright 2013 Matt Ho
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

require File.expand_path(File.dirname(__FILE__) + '/tasks/debian')

task :default => :package

task :clean => "go:clean" do
  exec "rm -rf #{DIST}"
  exec "rm -f *.deb"
end

namespace :go do 
  desc "go build"
  task :build do
    exec "(cd secret-tool ; go build)"
  end

  desc "go test"
  task :test do
    exec "go get -d -v ./..."
    exec "go test ./..."
  end

  desc "go clean"
  task :clean do
    exec "go clean ./..."
  end
end

namespace :vagrant do
  desc "vagrant up"
  task :up => "deb:package" do
    exec "vagrant up"
  end

  desc "vagrant destroy"
  task :destroy do
    exec "vagrant destroy"
  end

  desc "vagrant build"
  task :build do
    exec "BUILD=true vagrant up"
  end
end

namespace :deb do 
  desc "package"
  task :package => :prepare do 
    create_package
  end

  desc "prepare"
  task :prepare => "go:build" do 
    prepare_package
  end
end
