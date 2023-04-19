#!/usr/bin/env python3
# -*- coding=utf8 -*-

import sys
import delete_harbor_image
import pykube
import time

class GetK8sApi(object):
    # 判断harbor传入的项目是否存在k8s中
    def get_deployment(self, namespace, deployment_name):
        result = None
        try:
            api = pykube.HTTPClient(pykube.KubeConfig.from_file("k8s config"))
            deploy_out = pykube.Deployment.objects(api).filter(namespace=namespace).get(name=deployment_name)
            result = deploy_out.obj["spec"]["template"]["spec"]["containers"][0]["image"].split(":")[1]
            return result
        except pykube.exceptions.ObjectDoesNotExist:
            return result

class GetHarborApi(object):
    def __init__(self, host, user, password, protocol="http"):
        self.host = host
        self.user = user
        self.password = password
        self.protocol = protocol
        self.client = delete_harbor_image.HarborClient(self.host, self.user, self.password, self.protocol)

    # Del repo_name tag
    def del_repo_name_tag(self, repo_name, tag):
        for repo_name_tag in self.client.get_repository_tags(repo_name):
            if repo_name_tag['name'] != tag:
                self.client.del_repository_tag(repo_name, repo_name_tag['name'])

    # Del repo_name
    def del_repo_name(self, repo_name):
        self.client.delete_repository(repo_name)

    # 软删除harbor不用的项目及镜像
    def main(self):
        get_k8s_api = GetK8sApi()
        for projects in self.client.get_projects():
            # 公共基础镜像不用清理
            if projects['name'] != "public" and projects['name'] != "library" and len(
                    self.client.get_repositories(projects['project_id'])):
                for project_repo_name in self.client.get_repositories(projects['project_id']):
                    k8s_image_tag = get_k8s_api.get_deployment(project_repo_name['name'].split("/")[0], project_repo_name['name'].split("/")[1])
                    if k8s_image_tag is None:
                        self.del_repo_name(project_repo_name['name'])
                        time.sleep(3)
                    else:
                        self.del_repo_name_tag(project_repo_name['name'], k8s_image_tag)
                        time.sleep(3)
            # harbor中项目名和k8s中deploymnet名称不一致，暂时不处理，下面比如 test 项目
            elif projects['name'] == "test":
                pass

if __name__ == '__main__':
    sys.path.append("../")
    host = "domain"
    user = "user"
    password = "******"
    protocol = "https"
    cline_get = GetHarborApi(host, user, password, protocol)
    cline_get.main()