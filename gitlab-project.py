#!-*- encoding: utf-8 -*-
'''
Created on 2016年12月7日

@author: perry

使用python-gitlab模块对gitlab进行组添加的操作

https://python-gitlab.readthedocs.io/en/stable/index.html
http://gitlab.corp.lilyenglish.com/help/api/projects.md
'''
import gitlab
import time
import sys

# reload(sys)
# sys.setdefaultencoding("utf-8")
"""
gitlab.Project.VISIBILITY_INTERNAL 10
gitlab.Project.VISIBILITY_PRIVATE 0
gitlab.Project.VISIBILITY_PUBLIC 20
"""


class GitlabProject(object):

    def __init__(self):
        # self.gl = gitlab.Gitlab('http://git.corp.bianlifeng.com', 'Ame76pkKqnKfBGx2UJnJ')
        self.gl = gitlab.Gitlab('http://172.16.0.54/', 'bJ-Nm6BZCxG6n-aF8aVt')
        self.default_project_info = {
            'import_url': 'http://172.16.0.54/init/bare.git',
        }
        self.visibility = {"private": gitlab.VISIBILITY_PRIVATE,
                           "internal": gitlab.VISIBILITY_INTERNAL,
                           # too dangerous to open
                           # "public": gitlab.Project.VISIBILITY_PUBLIC,
                           }

    def __clear_data(self, data):
        """
        判断是list或者字符串，将数据做trim处理后返回
        """
        if isinstance(data, list):
            return filter(lambda x: x, [x.strip() for x in data])
        if isinstance(data, str):
            return data.strip()

    def __get_visiblility_level_id(self, visibility_level):
        """
        获取可见level
        """
        visibility_level = self.__clear_data(visibility_level).lower()
        return self.visibility.get(
            visibility_level,
            gitlab.VISIBILITY_INTERNAL)

    def validate_group(self, group_name):
        """
        验证是否存在该组，如果存在则返回group id
        """
        groups = self.gl.groups.list(all=True)
        for group in groups:
            if group.name == group_name:
                return True, group.id
        return False, -1

    def check_project_exists(self, group_id, project_name):
        """
        检查该组下是否已经存在同名工程

        Return
            已经存在则返回True
        """
        try:
            group = self.gl.groups.get(group_id)
            for prj in group.projects.list():
                if prj.name == project_name:
                    print("existed project, reject!!!!!!")
                    return True
            print("new project, have fun!!!")
            return False
        except:
            print("ERROR]] Meeting error with specific group id: %d" % group_id)
            return False

    def create_project(self, group_name, project_name, visibility_level_name, desc=""):
        """
        创建工程
        """
        group_name = self.__clear_data(group_name)
        project_name = self.__clear_data(project_name)
        visibility_level_name = self.__clear_data(visibility_level_name)
        desc = self.__clear_data(desc)

        is_exist, group_id = self.validate_group(group_name)
        if not is_exist:
            print("Not exist such group: %s" % group_name)
            sys.exit(1)

        if self.check_project_exists(group_id, project_name):
            print(
                "Project %s alreasy exists under %s, \
                please try another project name" %
                (project_name, group_name))
            sys.exit(1)

        visibility_level = self.__get_visiblility_level_id(
            visibility_level_name)

        # 定制需要创建的工程信息
        custom_project_info = {
            'name': project_name,
            'namespace_id': group_id,
            'visibility_level': visibility_level,
            'description': desc}
        self.default_project_info.update(custom_project_info)

        print(self.default_project_info)

        # 创建工程
        project = self.gl.projects.create(self.default_project_info)

        # project.pretty_print()

        # 等待创建工程完毕
        wait_count = 0
        find_master = False
        while (wait_count < 300):
            wait_count += 1
            try:
                branch = project.branches.get("master")
                branch.protect()
                find_master = True
                break
            except gitlab.exceptions.GitlabGetError:
                time.sleep(1)
        print("cost %.2f sec to wait creating project" % (wait_count * 1,))

        # 将工程的master分支设置为保护分支
        if not find_master:
            print("[ERROR] Fail to make master protected branch!")
        print(
            "Succeed to create project: http://gitlab.dayuan1997.com/%s/%s" %
            (group_name, project_name))


if __name__ == '__main__':
    if len(sys.argv) < 4:
        print("Usage: %s group_name project_name visible" % sys.argv[0])
        sys.exit(1)
    gp = GitlabProject()
    if len(sys.argv) >= 5:
        project_desc = " ".join(sys.argv[4:])
        print(sys.argv[1], sys.argv[2], sys.argv[3], project_desc)
        gp.create_project(sys.argv[1], sys.argv[2], sys.argv[3], project_desc)
    else:
        gp.create_project(sys.argv[1], sys.argv[2], sys.argv[3])

    # gp = GitlabProject()
    # gp.create_project("fe-labs", "porsche-android_T", "Internal", "加密库android端")
