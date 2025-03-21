import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import 'study_room_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int _selectedIndex = 0;
  String? _username;

  @override
  void initState() {
    super.initState();
    _checkLoginStatus();
  }

  Future<void> _checkLoginStatus() async {
    final prefs = await SharedPreferences.getInstance();
    final userData = prefs.getString('user_data');
    final sessionToken = prefs.getString('session_token');

    if (userData == null || sessionToken == null) {
      if (mounted) {
        Navigator.pushReplacementNamed(context, '/login');
        return;
      }
    }

    if (userData != null) {
      final user = jsonDecode(userData);
      setState(() {
        _username = user['username'];
      });
    }
  }

  final List<NavigationRailDestination> _destinations = const [
    NavigationRailDestination(
      icon: Icon(Icons.person_outline),
      selectedIcon: Icon(Icons.person),
      label: Text('用户'),
    ),
    NavigationRailDestination(
      icon: Icon(Icons.people_outline),
      selectedIcon: Icon(Icons.people),
      label: Text('好友'),
    ),
    NavigationRailDestination(
      icon: Icon(Icons.school_outlined),
      selectedIcon: Icon(Icons.school),
      label: Text('自习室'),
    ),
    NavigationRailDestination(
      icon: Icon(Icons.forum_outlined),
      selectedIcon: Icon(Icons.forum),
      label: Text('社区'),
    ),
  ];

  final List<Color> _destinationColors = [
    Colors.blue,
    Colors.green,
    Colors.orange,
    Colors.purple,
  ];

  Widget _buildContent() {
    if (_selectedIndex == 2) {
      return const StudyRoomPage();
    }
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(
            (_destinations[_selectedIndex].selectedIcon as Icon).icon,
            size: 100,
            color: _destinationColors[_selectedIndex],
          ),
          const SizedBox(height: 16),
          Text(
            _destinations[_selectedIndex].label.toString(),
            style: TextStyle(
              fontSize: 24,
              color: _destinationColors[_selectedIndex],
              fontWeight: FontWeight.bold,
            ),
          ),
          if (_selectedIndex == 0 && _username != null) ...[            
            const SizedBox(height: 16),
            Text(
              '欢迎回来，$_username',
              style: const TextStyle(fontSize: 18),
            ),
          ],
        ],
      ),
    );
  }

  Future<void> _handleLogout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('user_data');
    await prefs.remove('session_token');
    if (mounted) {
      Navigator.pushReplacementNamed(context, '/login');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Row(
        children: [
          NavigationRail(
            selectedIndex: _selectedIndex,
            onDestinationSelected: (int index) {
              setState(() {
                _selectedIndex = index;
              });
            },
            labelType: NavigationRailLabelType.all,
            destinations: _destinations,
            selectedIconTheme: IconThemeData(
              color: _destinationColors[_selectedIndex],
            ),
            selectedLabelTextStyle: TextStyle(
              color: _destinationColors[_selectedIndex],
            ),
            useIndicator: true,
            indicatorColor: _destinationColors[_selectedIndex].withOpacity(0.2),
            trailing: _username != null
                ? Padding(
                    padding: const EdgeInsets.only(bottom: 20),
                    child: TextButton.icon(
                      onPressed: _handleLogout,
                      icon: const Icon(Icons.logout),
                      label: const Text('登出'),
                      style: TextButton.styleFrom(
                        foregroundColor: Colors.red,
                      ),
                    ),
                  )
                : null,
          ),
          const VerticalDivider(thickness: 1, width: 1),
          Expanded(
            child: _buildContent(),
          ),
        ],
      ),
    );
  }
}