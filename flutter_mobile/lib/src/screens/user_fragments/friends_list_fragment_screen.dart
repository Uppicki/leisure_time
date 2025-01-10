import 'package:flutter/material.dart';

class FriendsListFragmentScreen extends StatelessWidget {
  const FriendsListFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Users"),
      ),
      body: ListView.builder(
        itemCount: 100,
        itemBuilder: (_, ind) => ListTile(
          title: Text("User $ind"),
          trailing: TextButton(
            onPressed: () {},
            child: Text("Write"),
          ),
        ),
      ),
    );
  }
}
