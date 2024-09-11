/** @type {import('tailwindcss').Config} */
export default {
  content: ["./web/*.html"],
  theme: {
    extend: {
      fontFamily: {
        "inter": ["Inter"]
      },
      colors: {
        "primary": "#6366f1",
        "secondary": "#e11d48",
        "on-primary": "#f4f4f5",
        "on-secondary": "#f5f5f4"
      }
    },
  },
  plugins: [],
}

