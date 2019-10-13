all:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	admincli combine js --path=./assets/src/js/combine/ --out=./assets/login/dist/all.min.js
	admincli combine css --path=./assets/src/css/ --out=./assets/login/dist/all.min.css
	admincli compile asset --path=./assets/login/dist/ --out=./ --pa=login