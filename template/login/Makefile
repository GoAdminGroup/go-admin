all:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	adm combine js --path=./assets/src/js/combine/ --out=./assets/login/dist/all.min.js
	adm combine css --path=./assets/src/css/ --out=./assets/login/dist/all.min.css
	adm compile asset --path=./assets/login/dist/ --out=./ --pa=login