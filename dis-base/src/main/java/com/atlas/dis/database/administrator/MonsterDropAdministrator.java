package com.atlas.dis.database.administrator;

import com.app.database.util.QueryAdministratorUtil;
import com.atlas.dis.entity.MonsterDrop;
import com.atlas.dis.model.MonsterDropData;

import javax.persistence.EntityManager;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public final class MonsterDropAdministrator {
   private MonsterDropAdministrator() {
   }

   public static void create(EntityManager entityManager, int monsterId, int itemId, int maximumQuantity, int minimumQuantity, int chance) {
      MonsterDrop monsterDrop = new MonsterDrop();
      monsterDrop.setMonsterId(monsterId);
      monsterDrop.setItemId(itemId);
      monsterDrop.setMaximumQuantity(maximumQuantity);
      monsterDrop.setMinimumQuantity(minimumQuantity);
      monsterDrop.setChance(chance);
      QueryAdministratorUtil.insert(entityManager, monsterDrop);
   }

   public static void createBulk(EntityManager entityManager, List<MonsterDropData> dropData) {
      QueryAdministratorUtil.insertBulk(entityManager,
            dropData.stream()
                  .filter(Objects::nonNull)
                  .map(data -> {
                     MonsterDrop monsterDrop = new MonsterDrop();
                     monsterDrop.setMonsterId(data.monsterId());
                     monsterDrop.setItemId(data.itemId());
                     monsterDrop.setMaximumQuantity(data.maximumQuantity());
                     monsterDrop.setMinimumQuantity(data.minimumQuantity());
                     monsterDrop.setChance(data.chance());
                     return monsterDrop;
                  })
                  .collect(Collectors.toList())
      );
   }
}
