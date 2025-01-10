import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';
import 'package:provider/provider.dart';

@RoutePage()
class ProfileFragmentScreen extends StatelessWidget {
  const ProfileFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final authService = context.read<AuthService>();

    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text("Profile"),
            FilledButton(
              onPressed: () {
                authService.logout();
              },
              child: Text("Logout"),
            ),
          ],
        ),
      ),
    );
  }
}
