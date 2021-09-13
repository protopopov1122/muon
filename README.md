### Muon CMS
Minimalistic Markdown-based CMS (well, it's too primitive to be CMS, but still). It's a custom
content engine that powers my [personal website](https://www.protopopov.lv), so repo includes not only
CMS engine, but also some specific website-related files.

Content engine is implemented in Golang with server-side page generation. No JavaScript is being used.
Backend is generating directly HTML by rendering Markdown documents. Previous version based on Python and VueJS
can be found in muon1 branch.

Engine is Dockerized. This repo does not contain fonts, so you should download them manually
into `data/static/` according to `data/static/style.css`.

#### License & Author
Author: Jevgēnijs Protopopovs

Project files are distributed under the terms of MIT License.
