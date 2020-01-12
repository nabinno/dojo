---
title: Deploy a Ruby on Rails App to App Engine Flexible Environment
tags: google-app-engine,ruby-on-rails,ruby
url: https://www.qwiklabs.com/focuses/2789?parent=catalog
---

# Goal
	- How to create a new Rails application using the Cloud Shell
- How to test a Rails application using Web Preview in the Cloud Shell before deployment
- How to add a Welcome page to the Rails application
- How to prepare your deployment configuration
- How to deploy a new Rails application to Google App Engine

# Task
- [x] Overview
- [x] Setup and Requirements
- [x] Enable the App Engine Admin API
- [x] Install the Rails gem
- [x] Generate a new Rails application
- [x] Test the Generated Rails Application
- [x] Add a Welcome page
- [x] Open a separate Cloud Shell tab
- [x] Deployment Configuration
- [x] Deploying the Application on App Engine
- [x] Check out your Rails App
- [x] Learn More
- [x] Test your knowledge
- [x] Congratulations!

# Supplement
## Install the Rails gem
```sh
gem install rails
rails --version
```

## Generate a new Rails application
```sh
rails new app_name
cd $_
ls
bin/bundle install
```

## Test the Generated Rails Application
```sh
bin/rails s -p 8080
```

## Add a Welcome page
```sh
bin/rails g controller Welcome index

cat <<EOF >app/views/welcome/index.html.erb
<h1>Welcome</h1>
<p>This is a home page for a new Rails App on Google Cloud Platform!</p>
EOF

cat <<EOF >config/routes.rb
Rails.application.routes.draw do
  get 'welcome/index'
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html

  root 'welcome#index'
end
EOF

bin/rails s -p 8080
```

## Deployment Configuration
```sh
cd app_name
export SECRET_KEY=$(bin/rails secret)
gcloud beta runtime-config configs create flex-env-config
gcloud beta runtime-config configs variables set \
  --config-name=flex-env-config \
  --is-text SECRET_KEY_BASE "${SECRET_KEY}"

cat <<EOF >app.yaml
# How to start your application
entrypoint: bundle exec rackup --port \$PORT
# Use App Engine flexible environment
env: flex
# Use the supported Ruby runtime
runtime: ruby

# App Engine flexible environment will load the configuration
# values from the Runtime Configuration Service defined by
# flex-env-config and create a .env file used by your
# application to determine environment values.
runtime_config:
  dotenv_config: flex-env-config
EOF

bin/bundle add dotenv-rails

export PROJECT_ID=$(gcloud config get-value project)
gcloud projects list --filter $PROJECT_ID
export PROJECT_NUMBER=$(gcloud projects list --filter $PROJECT_ID --format json | jq -r .[].projectNumber)
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:${PROJECT_NUMBER}@cloudbuild.gserviceaccount.com \
  --role=roles/editor
```

## Deploying the Application on App Engine
```sh
gcloud app create --region us-central
gcloud app deploy
```

## Check out your Rails App
```
https://${PROJECT_ID}.appspot.com
```

# References
- https://cloud.google.com/ruby/

