import 'package:auto_route/annotations.dart';
import 'package:flutter/material.dart';

@RoutePage()
class MapFragmentScreen extends StatelessWidget {
  const MapFragmentScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Map screen"),
      ),
      body: Center(
        child: Text("inProgress"),
      ),
    );
  }
}
