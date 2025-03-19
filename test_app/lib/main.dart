import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'routes/app_routes.dart';
import 'dart:async';
import 'dart:math';
import 'dart:ui' as ui;
import 'package:shared_preferences/shared_preferences.dart';
import 'dart:convert';
import 'package:webview_all/webview_all.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  void initState() {
    super.initState();
    _checkLoginStatus();
  }

  Future<void> _checkLoginStatus() async {
    final prefs = await SharedPreferences.getInstance();
    final sessionToken = prefs.getString('session_token');
    final userData = prefs.getString('user_data');

    if (sessionToken != null && userData != null) {
      // 如果存在有效的登录信息，导航到主页
      // AppRoutes.goToHome(navigatorKey.currentContext!);
    }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      //navigatorKey: navigatorKey,
      title: 'E - StudyRoom',
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFF10A37F),
          brightness: Brightness.light,
        ),
        textTheme: const TextTheme(
          displayLarge: TextStyle(fontSize: 32, fontWeight: FontWeight.bold),
          headlineMedium: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
          bodyLarge: TextStyle(fontSize: 16),
        ),
      ),
      routes: AppRoutes.routes,
      home: const HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  bool _showHoneycomb = false;
  String? _username;
  bool _isLoggedIn = false;

  @override
  void initState() {
    super.initState();
    _checkLoginStatus();
  }

  Future<void> _checkLoginStatus() async {
    final prefs = await SharedPreferences.getInstance();
    final userData = prefs.getString('user_data');
    if (userData != null) {
      final user = jsonDecode(userData);
      setState(() {
        _username = user['username'];
        _isLoggedIn = true;
      });
    }
  }

  Future<void> _handleLogout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('user_data');
    await prefs.remove('session_token');
    setState(() {
      _username = null;
      _isLoggedIn = false;
    });
  }

  void _onTypingFinished() {
    setState(() {
      _showHoneycomb = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        title: const Text(
          'E - StudyRoom',
          style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
        ),
        actions: [
          TextButton(
            onPressed: () {
              AppRoutes.navigateToApi(context);
            },
            child: const Text('API'),
          ),
          const SizedBox(width: 16),
          if (_isLoggedIn) ...[            
            Text(_username ?? '', style: const TextStyle(color: Colors.black54)),
            const SizedBox(width: 16),
            TextButton(
              onPressed: _handleLogout,
              child: const Text('Logout'),
            ),
          ] else
            TextButton(
              onPressed: () {
                AppRoutes.navigateToLogin(context);
              },
              child: const Text('Log in'),
            ),
          const SizedBox(width: 40),
        ],
      ),
      body: Stack(
        children: [
          //if (_showHoneycomb) HoneycombBackground(show: _showHoneycomb),
          SingleChildScrollView(
            child: Column(
              children: [
                Container(
                  height: MediaQuery.of(context).size.height - kToolbarHeight,
                  width: double.infinity,
                  decoration: BoxDecoration(
                    gradient: LinearGradient(
                      begin: Alignment.topCenter,
                      end: Alignment.bottomCenter,
                      colors: [
                        Theme.of(context).colorScheme.primary.withOpacity(0.1),
                        Theme.of(context).colorScheme.primary.withOpacity(0.2),
                        Theme.of(context).colorScheme.primary.withOpacity(0.1),
                      ],
                      stops: const [0.0, 0.5, 1.0],
                    ),
                  ),
                  child: Center(
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Container(
                          constraints: const BoxConstraints(maxWidth: 680),
                          child: TypewriterText(
                            text: 'Build Your Future, One Study Session at a Time.',
                            style: Theme.of(context).textTheme.headlineMedium!.copyWith(
                              color: Colors.black54,
                              fontWeight: FontWeight.normal,
                            ),
                            textAlign: TextAlign.center,
                            onTypingFinished: _onTypingFinished,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class TypewriterText extends StatefulWidget {
  final String text;
  final TextStyle? style;
  final TextAlign? textAlign;
  final VoidCallback? onTypingFinished;

  const TypewriterText({
    super.key,
    required this.text,
    this.style,
    this.textAlign,
    this.onTypingFinished,
  });

  @override
  State<TypewriterText> createState() => _TypewriterTextState();
}

class _TypewriterTextState extends State<TypewriterText> with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<double> _fadeAnimation;
  late Animation<double> _rotateAnimation;
  String _displayText = '';
  bool _showCursor = true;
  Timer? _cursorTimer;
  int _charIndex = 0;
  bool _typingFinished = false;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 2000),
    );

    _fadeAnimation = Tween<double>(
      begin: 1.0,
      end: 0.0,
    ).animate(CurvedAnimation(
      parent: _controller,
      curve: const Interval(0.0, 0.5, curve: Curves.easeOut),
    ));

    _rotateAnimation = Tween<double>(
      begin: 0.0,
      end: pi / 2,
    ).animate(CurvedAnimation(
      parent: _controller,
      curve: const Interval(0.0, 0.5, curve: Curves.easeOut),
    ));

    _startTyping();
    _startCursorBlink();
  }

  void _startTyping() {
    Timer.periodic(const Duration(milliseconds: 100), (timer) {
      if (_charIndex < widget.text.length) {
        setState(() {
          _displayText = widget.text.substring(0, _charIndex + 1);
          _charIndex++;
        });
      } else {
        timer.cancel();
        setState(() {
          _typingFinished = true;
        });
        widget.onTypingFinished?.call();
        _controller.forward();
      }
    });
  }

  void _startCursorBlink() {
    _cursorTimer = Timer.periodic(const Duration(milliseconds: 500), (timer) {
      if (mounted && !_typingFinished) {
        setState(() {
          _showCursor = !_showCursor;
        });
      }
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    _cursorTimer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Text.rich(
      TextSpan(
        children: [
          TextSpan(text: _displayText),
          if (!_typingFinished && _showCursor)
            TextSpan(
              text: '|',
              style: widget.style?.copyWith(
                color: Theme.of(context).colorScheme.primary,
              ),
            ),
        ],
      ),
      style: widget.style,
      textAlign: widget.textAlign,
    );
  }
}


// class HoneycombBackground extends StatefulWidget {
//   final bool show;
//   const HoneycombBackground({super.key, required this.show});

//   @override
//   State<HoneycombBackground> createState() => _HoneycombBackgroundState();
// }

// class _HoneycombBackgroundState extends State<HoneycombBackground> with SingleTickerProviderStateMixin {
//   late AnimationController _controller;
//   ui.FragmentShader? _shader;
//   double _time = 0.0;
//   bool _isLoading = true;
//   String? _error;

//   @override
//   void initState() {
//     super.initState();
//     _controller = AnimationController(
//       vsync: this,
//       duration: const Duration(seconds: 1),
//     )..repeat();

//     _loadShader();
//     if (widget.show) {
//       _controller.forward();
//     }
//   }

//   Future<void> _loadShader() async {
//     try {
//       final program = await ui.FragmentProgram.fromAsset('assets/shaders/honeycomb.frag');
//       if (mounted) {
//         setState(() {
//           _shader = program.fragmentShader();
//           _isLoading = false;
//         });
//       }
//     } catch (e) {
//       if (mounted) {
//         setState(() {
//           _error = e.toString();
//           _isLoading = false;
//         });
//       }
//       debugPrint('Error loading shader: $e');
//     }
//   }

//   @override
//   void didUpdateWidget(HoneycombBackground oldWidget) {
//     super.didUpdateWidget(oldWidget);
//     if (widget.show && !oldWidget.show) {
//       _controller.forward();
//     }
//   }

//   @override
//   void dispose() {
//     _controller.dispose();
//     super.dispose();
//   }

//   @override
//   Widget build(BuildContext context) {
//     if (_isLoading) {
//       return const Center(child: CircularProgressIndicator());
//     }

//     if (_error != null || _shader == null) {
//       return const SizedBox.expand();
//     }

//     return AnimatedBuilder(
//       animation: _controller,
//       builder: (context, child) {
//         _time += 0.016;
//         return CustomPaint(
//           painter: HoneycombShaderPainter(_shader!, _time),
//           size: Size.infinite,
//         );
//       },
//     );
//   }
// }

// class HoneycombShaderPainter extends CustomPainter {
//   final ui.FragmentShader shader;
//   final double time;

//   HoneycombShaderPainter(this.shader, this.time);

//   @override
//   void paint(Canvas canvas, Size size) {
//     shader.setFloat(0, time);
//     shader.setFloat(1, size.width);
//     shader.setFloat(2, size.height);

//     final paint = Paint()
//       ..shader = shader;

//     canvas.drawRect(
//       Rect.fromLTWH(0, 0, size.width, size.height),
//       paint,
//     );
//   }

//   @override
//   bool shouldRepaint(covariant HoneycombShaderPainter oldDelegate) {
//     return oldDelegate.time != time;
//   }
// }
