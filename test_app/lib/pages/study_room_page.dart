import 'package:flutter/material.dart';
import 'dart:math';
import 'package:flutter/services.dart';



class ZoomWrapper{
    static const platform = MethodChannel('win_zoom');


    static Future<String?> getPlatformVersion() async {
        try {
            final version = await platform.invokeMethod<String>('getPlatformVersion');
            return version;
        } on PlatformException catch (e) {
            print("Failed to get platform version: '${e.message}'.");
            return null;
        }
    }


    static Future<bool> initSDK() async {
        try {
            final result = await platform.invokeMethod<bool>('initSDK');
            return result ?? false;
        } on PlatformException catch (e) {
            print("Failed to initialize Zoom SDK: '${e.message}'.");
            return false;
        }
    }
}



class StudyRoomPage extends StatefulWidget {
    const StudyRoomPage({super.key});

    @override
    State<StudyRoomPage> createState() => _StudyRoomPageState();
}

class _StudyRoomPageState extends State<StudyRoomPage> {
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
                        child: GridView.builder(
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
                                        onTap: () async {
                                            if (index == 0) {
                                                final result = await ZoomWrapper.initSDK();
                                                if (mounted) {
                                                    ScaffoldMessenger.of(context).showSnackBar(
                                                        SnackBar(
                                                            content: Text(
                                                                result ? 'Zoom SDK 初始化成功' : 'Zoom SDK 初始化失败',
                                                            ),
                                                            backgroundColor: result ? Colors.green : Colors.red,
                                                        ),
                                                    );
                                                }
                                            }
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