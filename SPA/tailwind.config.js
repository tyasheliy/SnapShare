/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./public/**/*.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    colors: {
        'primary': '#1E1E24',
        'secondary': '#7D8491',
        'foreground': '#FFFBFE',
        'accent': '#32CBFF',
        'accept': '#2DD881',
        'error': '#FE5F55'
    },
    fontFamily: {
        inter: ['Inter', 'Helvetica', 'sans-serif']
    },
    extend: {},
  },
  plugins: [],
}

