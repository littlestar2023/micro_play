#FROM scratch
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/centos:centos7.9.2009

COPY ./task_service /opt/service/task_service
COPY ./config.yaml /opt/service/config.yaml

EXPOSE 8080

# push repository
# registry.cn-hangzhou.aliyuncs.com/wxg-uat/micro-task-service
CMD ["/opt/service/task_service"]