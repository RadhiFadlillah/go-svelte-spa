import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig(({ _, mode }) => {
	let buildConfig
	if (mode === 'production') {
		buildConfig = {
			emptyOutDir: true,
			outDir: '../build/public'
		}
	} else {
		buildConfig = {
			watch: {},
			minify: false,
			emptyOutDir: true,
			outDir: '../build/public-dev'
		}
	}

	return {
		root: 'frontend/',
		plugins: [svelte()],
		build: buildConfig
	}
})
