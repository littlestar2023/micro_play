FROM scratch

COPY ./user_service /opt/service/user_service
COPY ./config.yaml /opt/service/config.yaml

EXPOSE 8080

# push repository
# registry.cn-hangzhou.aliyuncs.com/wxg-uat/micro-user-service
CMD ["/opt/service/task_service"]