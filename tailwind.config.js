/** @type {import('tailwindcss').Config} */
export const content = [
	'./templates/**/*.templ',
	'./views/**/*.templ',
];
export const theme = {
	daisyui: {
		themes: ["black"],
	},
};

export const plugins = [
	require("daisyui"),
];
export const corePlugins = { preFlight: true };

