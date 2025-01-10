import 'package:flutter/material.dart';

class ProfileFragmentScreen extends StatelessWidget {
  const ProfileFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text("Profile"),
            FilledButton(
              onPressed: () {},
              child: Text("Logout"),
            ),
          ],
        ),
      ),
    );
  }
}
