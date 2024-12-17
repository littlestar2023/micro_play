FROM scratch

COPY ./task_service /opt/service/task_service
COPY ./config.yaml /opt/service/config.yaml

EXPOSE 8080

# push repository
# registry.cn-hangzhou.aliyuncs.com/wxg-uat/micro-task-service
CMD ["/opt/service/task_service"]