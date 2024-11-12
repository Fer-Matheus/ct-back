package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedDiffs() {
	diffs := []models.Diff{
		{

			Content: ": \"diff --git a/CHANGELOG.md b/CHANGELOG.md\nindex f0762d2a7..b7a7f412a 100644\n--- a/CHANGELOG.md\n+++ b/CHANGELOG.md\n@@ -1,4 +1,5 @@\n ## 0.9.3-rc2\n+ * STORM-558: change \"swap!\" to \"reset!\" to fix assignment-versions in supervisor\n  * STORM-555: Storm json response should set charset to UTF-8\n  * STORM-513: check heartbeat from multilang subprocess\n  * STORM-549: \"topology.enable.message.timeouts\" does nothing\", ",
		},
		{

			Content: ": \"diff --git a/src/main/java/com/google/devtools/build/lib/runtime/CommandEnvironment.java b/src/main/java/com/google/devtools/build/lib/runtime/CommandEnvironment.java\nindex 73c00d3fa2..1208995695 100644\n--- a/src/main/java/com/google/devtools/build/lib/runtime/CommandEnvironment.java\n+++ b/src/main/java/com/google/devtools/build/lib/runtime/CommandEnvironment.java\n@@ -166,10 +166,10 @@ public final class CommandEnvironment {\n   }\n \n   /**\n-   * Return an ordered version of the client environment restricted to those variables\n-   * whitelisted by the command-line options to be inheritable by actions.\n+   * Return an ordered version of the client environment restricted to those variables whitelisted\n+   * by the command-line options to be inheritable by actions.\n    */\n-  private Map<String, String> getCommandlineWhitelistedClientEnv() {\n+  public Map<String, String> getWhitelistedClientEnv() {\n     Map<String, String> visibleEnv = new TreeMap<>();\n     for (String var : visibleClientEnv) {\n       String value = clientEnv.get(var);\n@@ -426,7 +426,7 @@ public final class CommandEnvironment {\n         getCommandId(),\n         // TODO(bazel-team): this optimization disallows rule-specified additional dependencies\n         // on the client environment!\n-        getCommandlineWhitelistedClientEnv(),\n+        getWhitelistedClientEnv(),\n         timestampGranularityMonitor);\n   }\n \ndiff --git a/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoCommand.java b/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoCommand.java\nindex 0d7fb6abab..b5e596fb88 100644\n--- a/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoCommand.java\n+++ b/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoCommand.java\n@@ -193,29 +193,31 @@ public class InfoCommand implements BlazeCommand {\n \n   static Map<String, InfoItem> getHardwiredInfoItemMap(OptionsProvider commandOptions,\n       String productName) {\n-    List<InfoItem> hardwiredInfoItems = ImmutableList.<InfoItem>of(\n-        new InfoItem.WorkspaceInfoItem(),\n-        new InfoItem.InstallBaseInfoItem(),\n-        new InfoItem.OutputBaseInfoItem(productName),\n-        new InfoItem.ExecutionRootInfoItem(),\n-        new InfoItem.OutputPathInfoItem(),\n-        new InfoItem.BlazeBinInfoItem(productName),\n-        new InfoItem.BlazeGenfilesInfoItem(productName),\n-        new InfoItem.BlazeTestlogsInfoItem(productName),\n-        new InfoItem.CommandLogInfoItem(),\n-        new InfoItem.MessageLogInfoItem(),\n-        new InfoItem.ReleaseInfoItem(productName),\n-        new InfoItem.ServerPidInfoItem(productName),\n-        new InfoItem.PackagePathInfoItem(commandOptions),\n-        new InfoItem.UsedHeapSizeInfoItem(),\n-        new InfoItem.UsedHeapSizeAfterGcInfoItem(),\n-        new InfoItem.CommitedHeapSizeInfoItem(),\n-        new InfoItem.MaxHeapSizeInfoItem(),\n-        new InfoItem.GcTimeInfoItem(),\n-        new InfoItem.GcCountInfoItem(),\n-        new InfoItem.DefaultsPackageInfoItem(),\n-        new InfoItem.BuildLanguageInfoItem(),\n-        new InfoItem.DefaultPackagePathInfoItem(commandOptions));\n+    List<InfoItem> hardwiredInfoItems =\n+        ImmutableList.<InfoItem>of(\n+            new InfoItem.WorkspaceInfoItem(),\n+            new InfoItem.InstallBaseInfoItem(),\n+            new InfoItem.OutputBaseInfoItem(productName),\n+            new InfoItem.ExecutionRootInfoItem(),\n+            new InfoItem.OutputPathInfoItem(),\n+            new InfoItem.ClientEnv(),\n+            new InfoItem.BlazeBinInfoItem(productName),\n+            new InfoItem.BlazeGenfilesInfoItem(productName),\n+            new InfoItem.BlazeTestlogsInfoItem(productName),\n+            new InfoItem.CommandLogInfoItem(),\n+            new InfoItem.MessageLogInfoItem(),\n+            new InfoItem.ReleaseInfoItem(productName),\n+            new InfoItem.ServerPidInfoItem(productName),\n+            new InfoItem.PackagePathInfoItem(commandOptions),\n+            new InfoItem.UsedHeapSizeInfoItem(),\n+            new InfoItem.UsedHeapSizeAfterGcInfoItem(),\n+            new InfoItem.CommitedHeapSizeInfoItem(),\n+            new InfoItem.MaxHeapSizeInfoItem(),\n+            new InfoItem.GcTimeInfoItem(),\n+            new InfoItem.GcCountInfoItem(),\n+            new InfoItem.DefaultsPackageInfoItem(),\n+            new InfoItem.BuildLanguageInfoItem(),\n+            new InfoItem.DefaultPackagePathInfoItem(commandOptions));\n     ImmutableMap.Builder<String, InfoItem> result = new ImmutableMap.Builder<>();\n     for (InfoItem item : hardwiredInfoItems) {\n       result.put(item.getName(), item);\ndiff --git a/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoItem.java b/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoItem.java\nindex e8dc77f352..e8836e554f 100644\n--- a/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoItem.java\n+++ b/src/main/java/com/google/devtools/build/lib/runtime/commands/InfoItem.java\n@@ -46,6 +46,7 @@ import java.lang.management.ManagementFactory;\n import java.lang.management.MemoryMXBean;\n import java.lang.management.MemoryUsage;\n import java.util.Collection;\n+import java.util.Map;\n \n /**\n  * An item that is returned by <code>blaze info</code>.\n@@ -482,6 +483,29 @@ public abstract class InfoItem {\n     }\n   }\n \n+  /** Info item for the effective current client environment. */\n+  public static final class ClientEnv extends InfoItem {\n+    public ClientEnv() {\n+      super(\n+          \"client-env\",\n+          \"The specifications that need to be added to the project-specific rc file to freeze the\"\n+              + \" current client environment\",\n+          true);\n+    }\n+\n+    @Override\n+    public byte[] get(Supplier<BuildConfiguration> configurationSupplier, CommandEnvironment env)\n+        throws AbruptExitException {\n+      String result = \"\";\n+      for (Map.Entry<String, String> entry : env.getWhitelistedClientEnv().entrySet()) {\n+        // TODO(bazel-team): as the syntax of our rc-files does not support to express new-lines in\n+        // values, we produce syntax errors if the value of the entry contains a newline character.\n+        result += \"common --action_env=\" + entry.getKey() + \"=\" + entry.getValue() + \"\\n\";\n+      }\n+      return print(result);\n+    }\n+  }\n+\n   /**\n    * Info item for the default package. It is deprecated, it still works, when\n    * explicitly requested, but are not shown by default. It prints multi-line messages and thus\ndiff --git a/src/test/shell/integration/action_env_test.sh b/src/test/shell/integration/action_env_test.sh\nindex 017aa892d3..80976ed34b 100755\n--- a/src/test/shell/integration/action_env_test.sh\n+++ b/src/test/shell/integration/action_env_test.sh\n@@ -121,4 +121,26 @@ function test_latest_wins_env() {\n   expect_not_log \"FOO=foo\"\n }\n \n+function test_env_freezing() {",
		},
	}
	
	_ = database.Db.SaveDiffs(&diffs)
}