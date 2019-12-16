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
##
```sh

```

##
```sh

```

##
```sh

```

##
```sh

```



# 4. Case Study on Using Environments

