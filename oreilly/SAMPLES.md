# Archived exercise apps

The apps in this directory are exercise notes. They are not deployed, and they are not kept current.

`graphbook`, `hooks-news`, and `react_architecture_fundamentals` each had a lockfile, and two of them had both `package-lock.json` and `yarn.lock`. Dependabot then reported the same transitive advisory twice. Overrides on Create React App and webpack did not stay ahead of that tree.

Dependencies and lockfiles are removed. The source remains for reading. Do not run `npm install` or `yarn` here, and do not commit a lockfile.

Maintained dependency trees live in `datacamp/`, `aws_lambda/`, and `aws_cdk/cdk-workshop/`.
