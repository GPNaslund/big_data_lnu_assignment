import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
export default {
  kit: {
    adapter: adapter({
      pages: '../backend/internal/public',
      assets: '../backend/internal/public',
      fallback: 'index.html',
      precompress: false,
      strict: true
    }),
    paths: {
      base: ''
    }
  }
};
