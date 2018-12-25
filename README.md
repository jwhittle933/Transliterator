# Greek and Hebrew Transliterator
For tranliterating large chunks of Greek and Hebrew text (and evetually with support for Syriac)

## Tecnologies
Golang Echo framework on the backend, Vue for the UI. Backend uses template rendering to serve the UI, Vue instance written within `<script>` tags in the markup. In order to avoid clashes between `{{ }}` syntax in Echo and Vue, the best is to write ``{{ `{{ dataProperty }}` }}``.

## UI
![Image of UI](TransliteratorUI.png)