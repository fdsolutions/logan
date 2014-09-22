
# All tasks must be add under the namespace 'logan' 
namespace :logan do 

  desc "Run all test" 
  task :test do 
    sh "gom test"
  end


  namespace :gen do

    desc  'generate a spec file for go model'
    task :spec do
      model_name = ENV["MODEL_NAME"]
      if model_name
        sh "_vendor/bin/ginkgo generate #{model_name}"
      else
      puts "
          You must specify a model name using MODEL_NAME env variables.
          Eg: rake logan:generate:spec logan # => Create logan_test.go file.
        "
      end
    end
  end
  # Add all our tasks here ... 

end
