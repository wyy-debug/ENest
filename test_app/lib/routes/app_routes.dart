import 'package:flutter/material.dart';
import '../pages/api_page.dart';
import '../pages/login_page.dart';

class AppRoutes {
  // 路由名称常量
    static const String home = '/';
    static const String api = '/api';
    static const String login = '/login';

    // 路由配置表
    static Map<String, WidgetBuilder> routes = {
        api: (context) => const ApiPage(),
        login: (context) => const LoginPage(),
    };

    // 页面导航方法
    static void navigateToApi(BuildContext context) {
        Navigator.pushNamed(context, api);
    }

    static void navigateToLogin(BuildContext context) {
        Navigator.pushNamed(context, login);
    }

  // 返回上一页
    static void goBack(BuildContext context) {
        Navigator.pop(context);
    }
}