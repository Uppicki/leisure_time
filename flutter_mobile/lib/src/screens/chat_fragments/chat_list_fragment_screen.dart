import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/screens/chat_fragments/chat_fragment_screen.dart';

@RoutePage()
class ChatListFragmentScreen extends StatelessWidget {
  const ChatListFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Chats"),
      ),
      body: ListView.builder(
        itemCount: 5,
        itemBuilder: (context, currInd) => ListTile(
          title: Text("User $currInd"),
          subtitle: Text("Hello"),
          trailing: Text("$currInd"),
          onTap: () => Navigator.push(
            context,
            MaterialPageRoute(builder: (_) => ChatFragmentScreen()),
          ),
        ),
      ),
    );
  }
}
