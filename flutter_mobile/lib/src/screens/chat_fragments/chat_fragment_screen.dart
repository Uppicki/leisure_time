import 'package:flutter/material.dart';

class ChatFragmentScreen extends StatelessWidget {
  const ChatFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          onPressed: () => Navigator.pop(context),
          icon: Icon(Icons.arrow_back),
        ),
        title: Text("any chat"),
      ),
      body: ListView.builder(
        reverse: true,
        itemCount: 50,
        itemBuilder: (_, ind) => Align(
          alignment: ind.isEven ? Alignment.centerLeft : Alignment.centerRight,
          child: Container(
            margin: EdgeInsets.symmetric(vertical: 8.0),
            padding: EdgeInsets.all(16.0),
            decoration: BoxDecoration(
              color: ind.isEven ? Colors.blue.shade100 : Colors.green.shade100,
              border: Border.all(
                color: Colors.black,
                width: 2,
              ),
              borderRadius: BorderRadius.circular(8),
            ),
            child: Text(
              "Message $ind",
              style: TextStyle(fontSize: 16),
            ),
          ),
        ),
      ),
    );
  }
}
