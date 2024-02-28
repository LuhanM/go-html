# go-htmx

I created this repository to study about Golang with HTMX

https://www.youtube.com/watch?v=J3QXFClUSlc&list=PL8M9ZjrDX7lpqFj1NzRKoLgE1wGOMNRsi



tailwind docs
https://tailwindcss.com/docs
https://www.youtube.com/watch?v=ft30zcMlFao


* reading qrcode
    https://github.com/cozmo/jsQR?tab=readme-ov-file

1 - npm install jsqr --save


setup:
npm install -D tailwindcss
npx tailwindcss init


HTMX setup
===========================================================
> npm init
> npm install -D prettier prettier-plugin-go-template
> touch prettier.config.cjs

    const config = {
        plugins: ["prettier-plugin-go-template"],
        override: [
            {
                files: ["*.html"],
                option: {
                    parser: "go-template"
                },
            },
        ],
    };

    module.exports = config

> npm install -D vite
> npm install htmx.org

> touch vite.config.js 

    import { resolve } from "path";
    import { defineConfig } from "vite";

    export default defineConfig({
        build: {
            lib: {
                entry: [resolve(__dirname, "src/htmx.js")],
                formats: ["es"],
                name: "[name]",
                fileName: "[name]"
            },
            outDir: "static",
            emptyOutDir: false
        },
    });


- add `"type": "module",` inside package.json
- add "dev": "vite build --watch" in script section of package.json
- write src/htmx.js file
> npm run dev