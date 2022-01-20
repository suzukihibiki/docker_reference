FROM nginx:latest
COPY ./nginx.conf /etc/nginx/nginx.conf
# copyの左辺がローカルからコピーするもの
# copyの右辺が新しくdockerが作る仮想のディレクトリ
    # ので、docker上で何かやるときは右辺を指定する。