import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/app_router/app_router.gr.dart';

@RoutePage()
class AppNavScreen extends StatelessWidget {
  const AppNavScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return AutoTabsScaffold(
      routes: const [
        MapFragmentRoute(),
        SocialShellRoute(),
        ChatShellRoute(),
        ProfileFragmentRoute(),
      ],
      bottomNavigationBuilder: (_, tabsRouter) => BottomNavigationBar(
        currentIndex: tabsRouter.activeIndex,
        onTap: tabsRouter.setActiveIndex,
        items: const [
          BottomNavigationBarItem(
            icon: Icon(
              Icons.map,
              color: Colors.black,
            ),
            label: "1",
          ),
          BottomNavigationBarItem(
            icon: Icon(
              Icons.social_distance,
              color: Colors.black,
            ),
            label: "2",
          ),
          BottomNavigationBarItem(
            icon: Icon(
              Icons.chat,
              color: Colors.black,
            ),
            label: "3",
          ),
          BottomNavigationBarItem(
            icon: Icon(
              Icons.verified_user,
              color: Colors.black,
            ),
            label: "4",
          ),
        ],
      ),
    );
  }
}
