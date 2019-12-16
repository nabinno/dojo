---
title: Conda Essentials
tags: python,anaconda
url: https://www.datacamp.com/courses/conda-essentials
---

# 1. Installing Packages
## Install a specific version of a package (II)
```sh
conda install 'attrs>16,<17.3'
```

## Update a conda package
```sh
conda update pandas
```

## Remove a conda package
```sh
conda remove pandas
```

## Seach for available package versions?
```sh
conda search attrs
```

## Find dependencies for a package version?
```sh
conda search numpy=1.13.1 --info
```

# 2. Utilizing Channels
## Searching within channels
```sh
conda search -c conda-forge -c davidmertz -c sseefeld -c gbrener --platform osx-64 textadapter
```

## Searching across channels
```sh
anaconda search boltons
```

## Default, non-default, and special channels
```sh
conda search -c conda-forge | wc -l
```

## Installing from a channel
```sh
conda install -c conda-forge youtube-dl
conda list -c conda-forge
```

# 3. Working with Environments
## Which environment am I using?
```sh
conda env list
```

## What packages are installed in an environment? (I)
```sh
conda list 'numpy|pandas'
```

## What packages are installed in an environment? (II)
```sh
conda list --name test-env 'numpy|pandas'
```

## Switch between environments
```sh
conda env list
conda activate course-env
conda activate pd-2015
conda deactivate
```

## Remove an environment
```sh
conda env list
conda env remove --name deprecated
```

## Create a new environment
```sh
conda create --name conda-essentials attrs=19.1.0 cytoolz
conda activate conda-essentials
conda list
```

## Export an environment
```sh
conda env export -n course-env > course-env.yml
```

## Create an environment from a shared specification
```sh
cat environment.yml
conda env create --file environment.yml
conda env create --file shared-config.yml
```

# 4. Case Study on Using Environments
## Compatibility with different versions
```sh
cat weekly_humidity.py
python weekly_humidity.py
conda activate pd-2015
python weekly_humidity.py
```

## Updating a script
```sh
nano weekly_humidity.py
python weekly_humidity.py
conda activate pd-2015
python weekly_humidity.py
```
