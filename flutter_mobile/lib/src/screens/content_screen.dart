import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';

@RoutePage()
class ContentScreen extends StatelessWidget {
  const ContentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Text("isContent"),
      ),
    );
  }
}
