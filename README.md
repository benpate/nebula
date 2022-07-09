# nebula 🪐

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/nebula)
[![Build Status](https://img.shields.io/github/workflow/status/benpate/nebula/Go/main)](https://github.com/benpate/nebula/actions/workflows/go.yml)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/nebula.svg?style=flat-square)](https://codecov.io/gh/benpate/nebula)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/nebula?style=flat-square)](https://goreportcard.com/report/github.com/benpate/nebula)
[![Version](https://img.shields.io/github/v/release/benpate/nebula?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/nebula/releases)

## A loosely coupled library for building and editing HTML content

Nebula is a container format for rich HTML content, along with a reference library of HTML layouts and widgets, along with transactions to change items within the container.

Each nebula item has a view-only method along ith an editor method.  This lets developers create in-place editors for all kinds of applications.

Instead of storing content items in a natural, nested format, nebula puts each one in a single-dimensional array, which simplifies direct access to individual content items.

## Widget Library (in-progress)

* WYSIWYG content editor
* Image Uploads (with hooks into [mediaserver](https://github.com/benpate/mediaserver))
* OEmbed content for pictures, videos, and other rich objects
* more to come

## Layout Library (in progress)

* Row Layouts
* Responsive Columns (coming soon)
* Tab Control (coming soon)

## DO NOT USE

This project is a work-in-progress, and should NOT be used by ANYONE, for ANY PURPOSE, under ANY CIRCUMSTANCES.  It is NOT READY FOR PRIME TIME, and is essentially GUARANTEED to blow up your computer, send your cat into an infinite loop, and combine your hot and cold laundry into a single cycle.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 🤔
