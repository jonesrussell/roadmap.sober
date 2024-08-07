export default ({ mode }) => ({
  base: mode === 'production' ? '/roadmap.sober/' : '/',
  outDir: 'public/dist'
})
