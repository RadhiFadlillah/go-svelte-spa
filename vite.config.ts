import { defineConfig } from 'vite'
import type { BuildOptions } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
	let buildOptions: BuildOptions
	if (mode === 'production') {
		buildOptions = {
			emptyOutDir: true,
			outDir: 'build/public'
		}
	} else {
		buildOptions = {
			watch: {},
			minify: false,
			emptyOutDir: true,
			outDir: 'build/public-dev'
		}
	}

	return {
		plugins: [svelte()],
		build: buildOptions
	}
})
