(You can find example reference for this template [here](https://github.com/gtforge/arm_service/blob/dev/README.md))

# BookStore [![Build Status](https://travis-ci.com/gtforge/BookStore.svg?token=TRAVIS_TOKEN&branch=dev)](https://travis-ci.com/gtforge/BookStore) [![codecov](https://codecov.io/gh/gtforge/BookStore_service/branch/dev/graph/badge.svg?token=CODECOV_TOKEN)](https://codecov.io/gh/gtforge/BookStore)

## Service's Responsibility

\< think SRP: one thing the service does, and only it does it accross the whole system >

## Business Justification

\< there must be some to justify a new service. describe it >

## [Business Criticality](https://confluence.gtforge.com/display/AR/Business+Criticality)

\< TBD with an Architect >

## Dependencies

| Target | Channel |
| --- | --- |
| \<service name> | either HTTP / Rabbit MQ / Service Hub / ? |
| \<service name> | either HTTP / Rabbit MQ / Service Hub / ? |

## Owner

\< R&D group name >: \< at least two developer names here >

## Structure

\< key insight about how the service is built; main design decisions >

## Key Endpoints

\< desribe 2-3 key endpoint of the service, its purpose and how data looks like. this should give tangible insight about how the service contract >

If there are key scheduled tasks - explain them instead of some of the endpoints.

## Installation

```bash
./install.sh
```

\< what needed to be done to install it. If possible, don't write anything here and instead make sure install.sh is updated >

## Settings

* \< setting name > - \< meaning >
