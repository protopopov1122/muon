### Muon CMS
Minimalistic Markdown-based CMS (well, it's to primitive to be CMS, but still). It's a custom
content engine that powers my [personal website](https://www.protopopov.lv), so repo includes not only
CMS engine, but also some specific website-related files.

Content engine consists of SPA frontend using Vue and Python backend on Flask. Backend is serving
static files and rendering Markdown documents into HTML chunks, that are requested by frontend via
a simple set of REST services. Backend is capable of rendering separate files and whole directories.

Engine is completely Dockerized. This repo does not contain fonts, so you should download them manually
into `app/src/assets/fonts` according to `app/src/assets/css/fonts.css`.

#### License & Author
Author: Jevgēnijs Protopopovs

Project files are distributed under the terms of MIT License.
