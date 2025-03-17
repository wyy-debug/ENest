import 'package:flutter/material.dart';
import '../pages/api_page.dart';
import '../pages/login_page.dart';
import '../pages/home_page.dart';
import '../pages/study_room_page.dart';

class AppRoutes {
  // 路由名称常量
    static const String home = '/home';
    static const String api = '/api';
    static const String login = '/login';
    static const String studyRoom = '/study-room';

    // 路由配置表
    static Map<String, WidgetBuilder> routes = {
        home: (context) => const HomePage(),
        api: (context) => const ApiPage(),
        login: (context) => const LoginPage(),
        studyRoom: (context) => const StudyRoomPage(),
    };

    // 页面导航方法
    static void navigateToApi(BuildContext context) {
        Navigator.pushNamed(context, api);
    }

    static void navigateToLogin(BuildContext context) {
        Navigator.pushNamed(context, login);
    }

    static void goToHome(BuildContext context) {
        Navigator.pushReplacementNamed(context, home);
    }

  // 返回上一页
    static void goBack(BuildContext context) {
        Navigator.pop(context);
    }
}