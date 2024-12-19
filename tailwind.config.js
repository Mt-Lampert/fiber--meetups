/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html",
  ],
  theme: {
    extend: {
      fontFamily: {
        monserrat: ["Monserrat", "sans-serif"]
      }
    },
  },
  plugins: [
    require('daisyui'),
  ],
}

