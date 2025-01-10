import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/screens/chat_fragments/chat_fragment_screen.dart';
import 'package:flutter_mobile/src/screens/chat_fragments/chat_list_fragment_screen.dart';
import 'package:flutter_mobile/src/screens/map_fragments/map_fragment_screen.dart';
import 'package:flutter_mobile/src/screens/profile_fragments/profile_fragment_screen.dart';
import 'package:flutter_mobile/src/screens/user_fragments/friends_list_fragment_screen.dart';

class AppNavScreen extends StatefulWidget {
  const AppNavScreen({super.key});

  @override
  State<AppNavScreen> createState() => _AppNavScreenState();
}

class _AppNavScreenState extends State<AppNavScreen> {
  late int _currentIndex;

  final List<GlobalKey<NavigatorState>> _navigatorStates = [
    GlobalKey<NavigatorState>(),
    GlobalKey<NavigatorState>(),
    GlobalKey<NavigatorState>(),
    GlobalKey<NavigatorState>(),
  ];

  Widget _buildNav(int index) {
    return Navigator(
      key: _navigatorStates[index],
      onGenerateRoute: (settings) {
        switch (index) {
          case 0:
            return MaterialPageRoute(builder: (_) => MapFragmentScreen());
          case 1:
            return MaterialPageRoute(
                builder: (_) => FriendsListFragmentScreen());
          case 2:
            return MaterialPageRoute(builder: (_) => ChatListFragmentScreen());
          case 3:
            return MaterialPageRoute(builder: (_) => ProfileFragmentScreen());
          default:
            return MaterialPageRoute(builder: (_) => ProfileFragmentScreen());
        }
      },
    );
  }

  void _selectTab(int newIndex) {
    if (newIndex != _currentIndex) {
      _currentIndex = newIndex;
    }
    setState(() {});
  }

  @override
  void initState() {
    super.initState();
    _currentIndex = 0;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: IndexedStack(
        index: _currentIndex,
        children: [
          _buildNav(0),
          _buildNav(1),
          _buildNav(2),
          _buildNav(3),
        ],
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _currentIndex,
        onTap: _selectTab,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.map, color: Colors.black),
            label: "Map",
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.fireplace_rounded, color: Colors.black),
            label: "Friens",
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.chat_sharp, color: Colors.black),
            label: "Chats",
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.supervised_user_circle, color: Colors.black),
            label: "User",
          ),
        ],
      ),
    );
  }
}
