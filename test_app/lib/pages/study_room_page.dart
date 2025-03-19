import 'package:flutter/material.dart';
import 'dart:math';
import 'package:flutter/services.dart';
import 'package:desktop_webview_window/desktop_webview_window.dart';
import 'dart:convert';

class ZoomUIPage extends StatefulWidget {
    @override
    State<ZoomUIPage> createState() => _ZoomUIPageState();
}

class _ZoomUIPageState extends State<ZoomUIPage> {
    Webview? _webview;

    @override
    void initState() {
        super.initState();
        _initWebView();
    }

    Future<void> _initWebView() async {
        try {
            final webview = await WebviewWindow.create(
                configuration: CreateConfiguration(
                    windowHeight: 720,
                    windowWidth: 1280,
                    title: 'Zoom Video SDK UI Toolkit',
                    userDataFolderWindows: null,
                ),
            );
            setState(() {
                _webview = webview;
            });
            await webview.launch();
            await webview.loadHtml(_getHtmlContent());
        } catch (e) {
            print('WebView initialization error: $e');
        }
    }

    @override
    void dispose() {
        _webview?.close();
        super.dispose();
    }

    @override
    Widget build(BuildContext context) {
        return const SizedBox.shrink();
    }

    String _getHtmlContent() {
        return '''<!DOCTYPE html>
            <html>
            <head>
                <title>Zoom Video SDK UI Toolkit</title>
                <meta http-equiv="Content-Security-Policy" content="default-src * 'self' 'unsafe-inline' 'unsafe-eval' data: blob: wasm-unsafe-eval:;">
                <meta http-equiv="Cross-Origin-Opener-Policy" content="same-origin">
                <meta http-equiv="Cross-Origin-Embedder-Policy" content="require-corp">
                <script>
                    WebAssembly.compileStreaming = async (response) => {
                        const buffer = await response.arrayBuffer();
                        return WebAssembly.compile(buffer);
                    };
                </script>
                <script src="https://unpkg.com/@zoom/videosdk-ui-toolkit"></script>
            </head>
            <body>
                <div id="zoom-ui-toolkit"></div>
                <script>
                    window.onload = function() {
                        try {
                            const zoomUIToolkit = new ZoomVideoSDK.UIToolkit({
                                container: document.getElementById('zoom-ui-toolkit')
                            });
                            zoomUIToolkit.init();
                        } catch (error) {
                            console.error('Zoom UI Toolkit initialization error:', error);
                        }
                    };
                </script>
            </body>
            </html>''';
    }
}

class StudyRoomPage extends StatefulWidget {
    const StudyRoomPage({super.key});

    @override
    State<StudyRoomPage> createState() => _StudyRoomPageState();
}

class _StudyRoomPageState extends State<StudyRoomPage> {
    bool _showWebView = false;

    @override
    Widget build(BuildContext context) {
        return Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                    Text(
                        '自习室',
                        style: Theme.of(context).textTheme.headlineMedium!.copyWith(
                            color: Colors.orange,
                            fontWeight: FontWeight.bold,
                        ),
                    ),
                    const SizedBox(height: 24),
                    Expanded(
                        child: _showWebView
                            ? ZoomUIPage()
                            : GridView.builder(
                                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                                    crossAxisCount: 2,
                                    crossAxisSpacing: 16,
                                    mainAxisSpacing: 16,
                                    childAspectRatio: 1.5,
                                ),
                                itemCount: 4,
                                itemBuilder: (context, index) {
                                    return Card(
                                        elevation: 2,
                                        child: InkWell(
                                            onTap: () {
                                                setState(() {
                                                    _showWebView = true;
                                                });
                                            },
                                            child: Padding(
                                                padding: const EdgeInsets.all(16.0),
                                                child: Column(
                                                    mainAxisAlignment: MainAxisAlignment.center,
                                                    children: [
                                                        Icon(
                                                            Icons.school,
                                                            size: 48,
                                                            color: Colors.orange,
                                                        ),
                                                        const SizedBox(height: 8),
                                                        Text(
                                                            '自习室 ${index + 1}',
                                                            style: const TextStyle(
                                                                fontSize: 18,
                                                                fontWeight: FontWeight.bold,
                                                            ),
                                                        ),
                                                        const SizedBox(height: 4),
                                                        Text(
                                                            '当前人数: ${Random().nextInt(20)}',
                                                            style: TextStyle(
                                                                color: Colors.grey[600],
                                                            ),
                                                        ),
                                                    ],
                                                ),
                                            ),
                                        ),
                                    );
                                },
                            ),
                    ),
                ],
            ),
        );
    }
}