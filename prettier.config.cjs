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