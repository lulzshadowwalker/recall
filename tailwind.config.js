import daisyui from 'daisyui'

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./internal/template/**/*.{templ,go}"],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
  daisyui: {
    themes: ["lofi"],
  },
}

