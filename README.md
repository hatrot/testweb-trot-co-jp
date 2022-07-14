# testweb-trot-co-jp

## はじめに
本プロジェクトは GitHub Actions の動作確認と、
GAE デプロイの挙動確認用です。

メモ：
gcloud app deploy --quiet app.yaml --project testweb-trot-co-jp --version v1 --promote
gcloud app services set-traffic default --splits 20220711t200914=1 -q
gcloud app versions list | grep 20......t.........0.00 | sort -r | awk '3 < NR{print $2}' | xargs gcloud app versions delete
gcloud app versions list | grep 20......t......... | sort | head | awk '{print $2}' | xargs echo
gcloud app versions list | grep ' 1.00 ' | sort -r | awk 'NR==1{print $2}' | xargs echo