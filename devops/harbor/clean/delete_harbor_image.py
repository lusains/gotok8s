#!/usr/bin/env python3
# -*- coding=utf8 -*-

import json
import logging
import requests

logging.basicConfig(level=logging.INFO)


class HarborClient(object):
    def __init__(self, host, user, password, protocol="http"):
        self.host = host
        self.user = user
        self.password = password
        self.protocol = protocol

        self.session_id = self.login()

    # def __del__(self):
    #     self.logout()

    def login(self):
        login_data = requests.post('%s://%s/c/login' %
                                   (self.protocol, self.host),
                                   data={'principal': self.user,
                                         'password': self.password})

        if login_data.status_code == 200:
            session_id = login_data.cookies.get('sid')

            logging.debug("Successfully login, session id: {}".format(
                session_id))
            return session_id
        else:
            logging.error("Fail to login, please try again")
            return None

    def logout(self):
        requests.get('%s://%s/c/logout' % (self.protocol, self.host),
                     cookies={'sid': self.session_id})
        logging.debug("Successfully logout")

    # GET /projects
    def get_projects(self, project_name=None, is_public=None):
        # TODO: support parameter
        result = None
        path = '%s://%s/api/projects' % (self.protocol, self.host)
        response = requests.get(path,
                                cookies={'sid': self.session_id})
        if response.status_code == 200:
            result = response.json()
            logging.debug("Successfully get projects result: {}".format(
                result))
        else:
            logging.error("Fail to get projects result")
        return result

    # GET /repositories
    def get_repositories(self, project_id, query_string=None):
        # TODO: support parameter
        result = None
        path = '%s://%s/api/repositories?project_id=%s' % (
            self.protocol, self.host, project_id)
        response = requests.get(path,
                                cookies={'sid': self.session_id})
        if response.status_code == 200:
            result = response.json()
            logging.debug(
                "Successfully get repositories with id: {}, result: {}".format(
                    project_id, result))
        else:
            logging.error("Fail to get repositories result with id: {}".format(
                project_id))
        return result

    # DELETE /repositories
    def delete_repository(self, repo_name, tag=None):
        # TODO: support to check tag
        # TODO: return 200 but the repo is not deleted, need more test
        result = False
        path = '%s://%s/api/repositories/%s' % (self.protocol,
                                                          self.host, repo_name)
        response = requests.delete(path,
                                   cookies={'sid': self.session_id})
        if response.status_code == 200:
            result = True
            print("Delete {} successful!".format(repo_name))
            logging.debug("Successfully delete repository: {}".format(
                repo_name))
        else:
            logging.error("Fail to delete repository: {}".format(repo_name))
        return result

    # Get /repositories/{repo_name}/tags
    def get_repository_tags(self, repo_name):
        result = None
        path = '%s://%s/api/repositories/%s/tags' % (
            self.protocol, self.host, repo_name)
        response = requests.get(path,
                                cookies={'sid': self.session_id}, timeout=60)
        if response.status_code == 200:
            result = response.json()
            logging.debug(
                "Successfully get tag with repo name: {}, result: {}".format(
                    repo_name, result))
        else:
            logging.error("Fail to get tags with repo name: {}".format(
                repo_name))
        return result

    # Del /repositories/{repo_name}/tags/{tag}
    def del_repository_tag(self, repo_name, tag):
        result = False
        path = '%s://%s/api/repositories/%s/tags/%s' % (
            self.protocol, self.host, repo_name, tag)
        response = requests.delete(path,
                                   cookies={'sid': self.session_id})
        if response.status_code == 200:
            result = True
            print("Delete {} {} successful!".format(repo_name, tag))
            logging.debug(
                "Successfully delete repository repo_name: {}, tag: {}".format(
                    repo_name, tag))
        else:
            logging.error("Fail to delete repository repo_name: {}, tag: {}".format(
                repo_name, tag))
        return result