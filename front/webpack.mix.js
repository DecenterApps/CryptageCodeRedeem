const mix = require('webpack-mix').mix;

mix.sass('src/style.scss', 'dist/assets')
  // .copy('src/index.html', 'dist')
  .copyDirectory('src/assets', 'dist/assets');
