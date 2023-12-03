git init
git branch -M main
git remote add origin https://github.com/wit-Roman/translate-golang-main.git
git add *
git commit -a -m "add"
git push -u -f origin main

git pull -f --allow-unrelated-histories