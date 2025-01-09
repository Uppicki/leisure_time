import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/screens/auth_screen.dart';
import 'package:flutter_mobile/src/screens/content_screen.dart';

/// The Widget that configures your application.
class MyApp extends StatefulWidget {
  const MyApp({
    super.key,
  });

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  late bool isAuth;

  @override
  void initState() {
    super.initState();
    isAuth = false;
  }

  void login() {
    isAuth = true;
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: !isAuth
          ? AuthScreen(
              callback: login,
            )
          : ContentScreen(),
    );
  }
}
