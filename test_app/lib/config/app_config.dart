class AppConfig {
  static const String apiBaseUrl = 'http://localhost:3000/api';

  // 身份验证相关的API端点
  static String get loginEndpoint => '$apiBaseUrl/auth/login';
  static String get registerEndpoint => '$apiBaseUrl/auth/register';
}